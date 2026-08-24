import { execFileSync } from 'node:child_process';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

// Seeds a fixed, known set of users into the isolated `driveaway_e2e`
// database (inside the db_test container) before the suite runs. Requires
// the backend to have started at least once against driveaway_e2e already
// (so AutoMigrate has created the schema) — see e2e/README.md for the full
// run order. Never touches the dev `db` or production RDS.
const __dirname = path.dirname(fileURLToPath(import.meta.url));
const REPO_ROOT = path.resolve(__dirname, '..');

const SEED_SQL = `
TRUNCATE TABLE trylogin, doubleregistration, personalinformation, phone, mail,
  name, password, users, sort, driveofdriver, driveavailable, driveravailable,
  drive, driver, permissions, company RESTART IDENTITY CASCADE;

INSERT INTO users (id, "userName") VALUES
  (1, 'e2e-admin'),
  (2, 'e2e-manager'),
  (3, 'e2e-driver');

INSERT INTO password (userid, password) VALUES
  (1, 'e2e-admin-pass123'),
  (2, 'e2e-manager-pass123'),
  (3, 'e2e-driver-pass123');

INSERT INTO company (id, name, idcompany, address) VALUES (1, 'E2E Test Co', 'E2E-1', 'Test Address 1');

INSERT INTO permissions (userid, companyid, level) VALUES
  (1, 1, 1),
  (2, 1, 3),
  (3, 1, 5);

-- One drive and one driver (the seeded e2e-driver user), both marked
-- available, so sort-creation.spec.ts has something to actually sort.
INSERT INTO drive (id, name, htostart, mtostart, gpsstart, gpsend, companyid) VALUES
  (1, 'E2E Base Drive', 8, 0, '31.7,34.8', '31.8,34.9', 1);
INSERT INTO driver (id, userid, companyid, address, isemployee) VALUES
  (1, 3, 1, 'Test Address 1', true);
INSERT INTO driveavailable (idcompany, iddrive, available) VALUES (1, 1, true);
INSERT INTO driveravailable (idcompany, iddriver, available) VALUES (1, 3, true);

-- Explicit ids above bypass each table's serial sequence, so the sequence's
-- internal counter is left at its post-TRUNCATE start value. The app's own
-- inserts (e.g. adding a driver/drive through the UI) don't specify an id
-- and rely on the sequence — without advancing it here, the very next
-- app-created row collides with one of the ids we just inserted, GORM's
-- Create() silently swallows that constraint-violation error (a real,
-- separate bug in the model layer worth its own fix later), and the row
-- ends up referencing id 0. Advancing every sequence we touched avoids
-- ever hitting that path in E2E runs.
SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));
SELECT setval('company_id_seq', (SELECT MAX(id) FROM company));
SELECT setval('drive_id_seq', (SELECT MAX(id) FROM drive));
SELECT setval('driver_id_seq', (SELECT MAX(id) FROM driver));
`;

export default function globalSetup() {
  execFileSync(
    'docker',
    ['compose', 'exec', '-T', 'db_test', 'psql', '-U', 'postgres', '-d', 'driveaway_e2e', '-v', 'ON_ERROR_STOP=1'],
    { cwd: REPO_ROOT, input: SEED_SQL, stdio: ['pipe', 'inherit', 'inherit'] }
  );
}
