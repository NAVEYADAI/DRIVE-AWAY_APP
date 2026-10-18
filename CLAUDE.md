# DRIVE-AWAY_APP

מערכת לניהול חברות הסעה: מנהל מערכת רושם חברות, מנהל חברה מוסיף נהגים ונסיעות, ואלגוריתם שיבוץ
(Mapbox) מחלק נסיעות לנהגים. כל נהג רואה את הנסיעות שלו.

מונוריפו: `back/` (Go, Gin + GORM + PostgreSQL, standard layout תחת `cmd/` + `internal/`) ו-`font/`
(React + TypeScript + Vite + MUI; שם התיקייה הוא באמת `font`). בנוסף, `e2e/` (Playwright). הפיתוח
המקומי רץ דרך `docker-compose.yml` (שירותים: `db`, `db_test`, `backend`, `frontend`).

## לפני שעובדים בתוך תת-תיקייה

כל תת-תיקייה מחזיקה `CLAUDE.md` → `AGENTS.md` משלה, שנטען אוטומטית כשעובדים על קבצים בתוכה. אל
תשכפל את התוכן שלהם כאן ואל תסתור אותו.

- **`back/AGENTS.md`**: מוסכמות Go/GORM, מודל ההרשאות `{userId, level}`, מלכודות בצורת התשובות,
  `testutil`, והכלל של AutoMigrate מול ה-DB של production.
- **`font/AGENTS.md`**: routes ו-HOCs של הרשאות, `GlobalState`, קריאות axios, `theme.ts`, ה-harness
  של Vitest.
- **`e2e/AGENTS.md`**: הכללים. `e2e/README.md` הוא סדר ההרצה המדויק.
- **`specs/`**: תיעוד עובדתי לכל תחום, נגזר מקריאת הקוד (`00` מוסכמות, `01` auth, `02` management,
  `03` sorting). **קרא את הקובץ הרלוונטי לפני שכותבים טסטים או נוגעים בהתנהגות קיימת.** הוא מקור האמת
  למקרי קצה ולבאגים ידועים (§9 בכל קובץ), ולא הנחה על מה שהקוד "אמור" לעשות.
- **`WORKFLOW.md`**: תוכנית העבודה (סדר שלבים), שערי Ready/Done, שיטות עבודה (branches, WIP=1,
  אדום→ירוק, שינויי סכמה בהוספה בלבד), וטבלת תקלות מוכרות. **קרא לפני שמתחילים משימה חדשה.**
- **`product-backlog/`**: באק לוג מוצרי, קובץ לכל פיצ'ר (תבנית ואינדקס ב-`README.md`). רעיונות
  שעוד לא אופיינו נמצאים ב-`IDEAS.md`.

## פקודות נפוצות (מהשורש)

- `docker compose up -d db_test`: DB הבדיקות המבודד (localhost:5434). נדרש לטסטים של Go.
- `cd back && go test ./...`: כל טסטי ה-backend.
- `cd font && npm test`: Vitest. `cd font && npx tsc -b`: בדיקת טיפוסים.
- `npm --prefix e2e test`: Playwright, רק אחרי שלבי ההקמה ב-`e2e/README.md`.
- `docker compose up -d`: סביבת פיתוח מלאה (frontend :3100 כברירת מחדל, אפשר לשנות עם `FRONTEND_PORT`; backend :2909).

## ⚠️ Git ו-deploy

- **push ל-`main` = deploy ל-AWS production** (`.github/workflows/deploy.yml`: build → ECR → ECS).
  אין CI שמריץ טסטים לפני ה-deploy. לכן לא עושים `git push` בלי אישור מפורש (הוא גם מוגדר כ-`ask`
  ב-`.claude/settings.json`), ועובדים על branch כשמשנים קוד.
- `back/.env` מצביע על ה-RDS של **production**. אסור לקרוא, להדפיס או להתחבר אליו (חסום
  ב-`.claude/settings.json`). גם `back/task-definition-backend.json` ו-`font/task-definition-frontend.json`
  (ב-gitignore) דורשים אישור לפני קריאה.
- שינוי סכמה עובר דרך `AutoMigrate` בעליית השרת, כלומר גם על production ב-deploy הבא. שינוי שם, מחיקה
  או שינוי טיפוס של עמודה קיימת דורשים אישור מפורש מנוה.

## מדיניות בדיקות (אומצה 2026-08-24)

הכללים האלה מחייבים. הם חלק מההגדרה של "סיימתי", לא עבודת המשך.

1. **כל תיקון באג מגיע עם טסט רגרסיה.** תיקון לא גמור עד שיש טסט שמשחזר בדיוק את מסלול הכשל ונכשל
   בלי התיקון. כשאפשר: לכתוב את הטסט קודם, לוודא שהוא נכשל, לתקן, ולוודא שהוא עובר. השם או הערה
   בשורה אחת צריכים לאפשר לעקוב אחרי הסימפטום המקורי (למשל `TestLogin_EmptyPasswordDoesNotCrash` עם
   הערה על מה נשבר ולמה). אם הבאג מתועד ב-`specs/*.md` §9, מעדכנים את השורה (STATUS → `Fixed <date>`
   + הפניה לטסט).
2. **Backend (`back/`)**: `testing` + `testify` (`require`/`assert`). הטסטים יושבים ליד הקוד. הקונבנציה
   בפועל היא **external test package** (`model_test`, `manager_test`, `sorting_test`...), ו-package
   פנימי רק כשצריך גישה ל-unexported (כמו `sorting/MapboxApi_test.go`). טסט שנוגע ב-DB מתחיל
   ב-`testutil.SetupTestDB(t)` (`back/internal/testutil`), ולעולם לא קורא ל-`config.Connect()`.
3. **Frontend (`font/`)**: Vitest + React Testing Library, `*.test.tsx` ליד הקומפוננטה. בודקים
   התנהגות שהמשתמש רואה (מה מוצג, מה לחיצה עושה) ולא פרטי מימוש.
4. **E2E (`e2e/`)**: Playwright מול ה-stack האמיתי (frontend + backend + `db_test`) לזרימות הקריטיות
   (login, CRUD מרכזי, יצירת שיבוץ). זו ראש הפירמידה: golden paths ומקרי הקצה החשובים ביותר, לא כל
   הפרמוטציות.
5. **בידוד DB**: `db_test` (פורט 5434, volume נפרד) מיועד רק לבדיקות, נפרד לגמרי מ-`db` ומה-RDS
   של production. טסטי Go משתמשים ב-DB `driveaway_test`, ו-E2E משתמש ב-`driveaway_e2e` על אותו שרת.
6. לפני שמכריזים שעבודה הסתיימה: backend → `cd back && go test ./...`; frontend → `cd font && npm test`
   (+ `npx tsc -b`); שינוי חוצה stack → גם ה-spec הרלוונטי ב-`e2e/`.

## הגדרות Claude Code בריפו (`.claude/`)

- **`settings.json`** (משותף, נכנס ל-git): הרשאות. go/npm test ו-git read-only מותרים. `git push`,
  `aws` ו-`docker compose down` דורשים אישור. קריאת `.env` ו-force-push חסומים. בנוסף hook שמריץ
  `gofmt -w` על כל קובץ `.go` שנערך. הגדרות אישיות נכנסות ל-`.claude/settings.local.json`
  (ב-gitignore).
- **הצוות (plugin `dev-team`)** — הסוכנים (`dev-team:tech-lead`, `product-manager`, `backend-developer`, `frontend-developer`,
  `qa-engineer`, `code-reviewer`, `security-checker`, `ux-ui-designer`, `dependency-checker`) מגיעים מ-plugin משותף
  (ריפו נפרד, מותקן ברמת משתמש) ולא יושבים כאן. **`TEAM.md` בשורש** הוא ה-adapter שהם קוראים ראשון: סטאק, קבצים לכל תפקיד,
  פקודות, כללים קשיחים, מילות הבאקלוג והגדרת "גמור". כללים מעדכנים ב-`TEAM.md`, לא בתוך סוכן. נקודת הכניסה לקידום
  הבאק לוג: `dev-team:tech-lead`.

## הערות ארכיטקטורה

- Backend: `back/cmd/server` (entrypoint, פורט 2909), `back/internal/{config,auth,management,sorting,testutil}`.
  חיבור ה-DB ב-`back/internal/config/db.go` נקבע מ-env (`DB_HOST`/`DB_PORT`/`DB_USER`/`DB_PASSWORD`/`DB_DBNAME`/`DB_SSLMODE`).
- **אין אימות בצד השרת.** הלקוח שולח `userId`/`level` בכל בקשה והשרת סומך עליהם. ר'
  `specs/00-shared-conventions.md` §3. לא מציגים endpoint חדש כ"מאובטח".
- Frontend: `font/src/theme.ts` הוא מקור יחיד לטוקנים של עיצוב (palette, shape, typography). לא
  מקודדים צבעים או מרווחים בקומפוננטות כשיש טוקן.
- CORS ב-backend מתיר כל `http://localhost:<port>` (`back/cmd/server/main.go`), כי פורט ה-dev
  server של ה-frontend יכול לזוז.
