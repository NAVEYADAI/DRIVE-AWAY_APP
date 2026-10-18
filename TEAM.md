# TEAM.md — DRIVE-AWAY_APP

> הקובץ ש-agents של ה-plugin `dev-team` קוראים ראשון. הוא מתאר איך הפרויקט **הזה** עובד, ומצביע לקבצים
> שמחזיקים את הכללים המפורטים (`CLAUDE.md`, `AGENTS.md`, `specs/`, `WORKFLOW.md`) — לא משכפל אותם. אם משהו
> כאן סותר את הקוד, הקוד מנצח ומתקנים את הקובץ.

## Project
מערכת לניהול חברות הסעה: מנהל מערכת רושם חברות, מנהל חברה מוסיף נהגים ונסיעות, ואלגוריתם שיבוץ (Mapbox)
מחלק נסיעות לנהגים. כל נהג רואה את הנסיעות שלו.

## Stack & layout
| Area | Path | Tech |
|---|---|---|
| backend | `back/` | Go, Gin + GORM + PostgreSQL; `cmd/server` (פורט 2909) + `internal/{config,auth,management,sorting,testutil}` |
| frontend | `font/` (כן, "font") | React + TypeScript + Vite + MUI; Vitest + React Testing Library |
| e2e | `e2e/` | Playwright מול stack אמיתי (frontend + backend + `db_test`) |
| specs | `specs/` | תיעוד עובדתי per-domain (`00` מוסכמות, `01` auth, `02` management, `03` sorting); §9 בכל קובץ = באגים ידועים |
| backlog | `product-backlog/` | קובץ לכל פיצ'ר + `README.md` + `IDEAS.md` |
| process | `WORKFLOW.md` | סדר שלבים, שערי Ready/Done, branches, WIP=1, אדום→ירוק |

פיתוח מקומי: `docker-compose.yml` (שירותים `db`, `db_test`, `backend`, `frontend`). פריסה: push ל-`main` = deploy ל-AWS
(ECR → ECS) **בלי CI שמריץ טסטים**.

## Conventions (read before working)
| Role | Files to read first |
|---|---|
| backend-developer | `back/AGENTS.md`, `specs/00-shared-conventions.md` + הספק של התחום (`01-auth`, `02-management`, `03-sorting`) |
| frontend-developer | `font/AGENTS.md` (+ `font/src/theme.ts` — טוקנים בלבד) |
| qa-engineer | `CLAUDE.md` (מדיניות בדיקות), `e2e/AGENTS.md`, `e2e/README.md` (סדר הרצה מדויק + fixtures), `back/AGENTS.md`, `font/AGENTS.md` |
| ux-ui-designer | `font/src/theme.ts`, `font/AGENTS.md` |
| security-checker | `specs/00-shared-conventions.md` §3, `specs/*.md` §9, `back/AGENTS.md` |
| code-reviewer | כל הנ"ל שרלוונטי ל-diff + `WORKFLOW.md` §2 (שערי Done) |
| product-manager | `product-backlog/README.md`, `product-backlog/IDEAS.md`, `specs/` |

## Commands
| Purpose | Command |
|---|---|
| DB בדיקות מבודד (localhost:5434) | `docker compose up -d db_test` |
| backend tests | `cd back && go test ./...` ; `cd back && go vet ./...` |
| frontend tests | `cd font && npm test` |
| typecheck | `cd font && npx tsc -b` |
| e2e | `npm --prefix e2e test` — **רק** אחרי שלבי ההקמה ב-`e2e/README.md`, ובסיום מחזירים את ה-backend ל-DB של dev (שלב אחרון שם) |
| dev stack | `docker compose up -d` (frontend :3100, backend :2909) |

## Hard rules (never)
- **`git push` — לעולם לא** (push ל-`main` = deploy ל-production; גם מוגדר `ask`). עובדים על branch.
- **`back/.env` מצביע על RDS של production** — אסור לקרוא, להדפיס או להתחבר (חסום ב-`.claude/settings.json`). גם
  `back/task-definition-backend.json` ו-`font/task-definition-frontend.json` דורשים אישור לפני קריאה.
- להתחבר **רק** ל-`db_test` (DB `driveaway_test` לטסטי Go, `driveaway_e2e` ל-E2E). לא ל-`db` ולא ל-RDS.
- שינוי סכמה עובר `AutoMigrate` בעליית השרת, כלומר גם על production ב-deploy הבא: הוספת טבלה/עמודה — בסדר; שינוי שם,
  מחיקה או שינוי טיפוס של עמודה קיימת → **עוצרים ומדווחים** לאישור מפורש של נוה.
- לא נוגעים ב-`.github/workflows/` בלי אישור. לא `docker compose down`, לא `aws`.

## Backlog
- dir: `product-backlog/` · ideas: `IDEAS.md`
- priority marker: `עדיפות מיידית`
- draft marker: `— טיוטה, ממתין להחלטות` (בכותרת: `## פיצ'ר N: <name> — טיוטה, ממתין להחלטות`)
- headings: Backend=`### N.1 Backend` · Frontend=`### N.2 Frontend` · Tests=`### N.3 בדיקות` ·
  Touchpoints=`### נקודות מרכזיות (touchpoints)`
- status strings: specified=`מאופיין במלואו, מוכן למימוש — טרם הותחל` · draft=`טיוטה — ממתין להחלטות (X שאלות פתוחות)` ·
  backend done=`Backend מומש, Frontend/בדיקות טרם הותחלו` · backend+frontend done=`Backend+Frontend מומש, בדיקות טרם הותחלו` ·
  done=`מומש ✅`
- **באגים:** אין `BUGS.md`. באגים ידועים ב-`specs/*.md` §9 ("Known Discrepancies & Risks"). תיקון באג הוא לא פיצ'ר; כשמתקנים,
  מעדכנים את שורת §9 (STATUS → `Fixed <date>` + הפניה לטסט הרגרסיה).

## Definition of done (extra gates)
- **כל תיקון באג מגיע עם טסט רגרסיה** שמשחזר את מסלול הכשל ונכשל בלי התיקון (`CLAUDE.md` §מדיניות בדיקות).
- backend → `cd back && go test ./...`; frontend → `cd font && npm test` + `npx tsc -b`; שינוי חוצה stack → גם ה-spec
  הרלוונטי ב-`e2e/`.
- `specs/*.md` מתעדכן כשהתנהגות מתועדת משתנה (backend: ה-spec של התחום; frontend: §5/§6).
- שורת הסטטוס ב-`product-backlog/README.md` מעודכנת; פיצ'ר "גמור" רק כשכל ה-`- [ ]` מסומנים (`WORKFLOW.md` §2).
- סקירת קוד (`dev-team:code-reviewer`) על ה-diff; הממצאים מטופלים או נדחים במפורש.

## Language & tone
- דיווחים למשתמש: **עברית**, תמציתי. טקסט UI: עברית (בקוד, `setNameOfPage("<עברית>")`).

## Role notes
### product
- בודקים מה כבר מגיע למסך (`GlobalState` מחזיק רק `level`, `userId`, `nameOfPage`) לפני שמניחים שדה API חדש.
- אם הרעיון נוגע בשורת §9 ("Bug candidate") — לציין אם הפיצ'ר תלוי בתיקונה.
- מחקר מודלים: `back/internal/*/model/*.go`, `handler/**`. דפוס לשימוש חוזר נקרא בשם (למשל "אותו gate ובדיקת בעלות כמו
  `manager.Delete_Drive`").
- ה-checklist של "בדיקות" מפרט את הכיסוי הנדרש לפי מדיניות הבדיקות: Go, Vitest, Playwright.

### backend
- מודל חדש נכנס לרשימת `AutoMigrate` של התחום **וגם** ל-`internal/testutil/db.go` (`AutoMigrate` + `allTables`).
- endpoint מנהל חדש משתמש באותו gate `{userId, level}` כמו אחיו; endpoint נהג **לא** משתמש ב-`ReturnPermissionForSingUp`.
- טסטים: `testing` + `testify`, package חיצוני (`<pkg>_test`); טסט שנוגע ב-DB מתחיל ב-`testutil.SetupTestDB(t)` ולעולם לא
  קורא ל-`config.Connect()`. handler tests דרך `httptest` + Gin router חשוף.
- אין אימות בצד השרת — הלקוח שולח `userId`/`level` והשרת סומך עליהם (`specs/00` §3). לא מציגים endpoint חדש כ"מאובטח".
- Go מחזיר `null` ולא `[]` לרשימה ריקה; מפתחות JSON כוללים שגיאות כתיב מתועדות (`aviable`, `messege`) — לא "מתקנים" בשקט.

### frontend
- מסך חדש: default export עטוף ב-HOC ההרשאה הנכון, `<Route>` ב-`src/App.tsx`, אפקט `setNameOfPage("<עברית>")`,
  נקודת כניסה (ActionCard במסך הבית של התפקיד או פריט תפריט), פריסת `Paper` סטנדרטית. צבע/מרווח מ-`theme.ts` בלבד.
- קריאות axios ל-`http://localhost:2909` עם `{..., userId, level}`; מתאימים בדיוק את מפתחות ה-JSON של ה-backend
  (כולל שגיאות הכתיב) ומגנים מפני `null` במקום `[]`.
- כל קומפוננטה חדשה/ששונתה מקבלת `*.test.tsx` צמוד (מצבי רינדור + payload של axios + טקסט alert, שאילתה לפי תווית עברית).

### qa
- Playwright: `e2e/tests/<feature>.spec.ts`. fixtures חדשים נכנסים ל-`SEED_SQL` ב-`e2e/global-setup.ts` ולא נוצרים ad hoc.
  מקבלים `alert` עם `page.on('dialog', ...)`. תמיד case של הרשאה (תפקיד שלא אמור להגיע — redirected/absent, לא רק מוסתר).
- משחזרים את מצב הסביבה בסיום (ה-backend חוזר ל-DB של dev לפי `e2e/README.md`).

### design
- `font/src/theme.ts` הוא מקור יחיד לטוקנים (palette, shape, typography). אין צבעים/מרווחים מקודדים כשיש טוקן.
- אין עדיין מדריך עיצוב נפרד או backlog מיגרציה — TODO: אם יוגדר, להוסיף כאן.

### security
- הסיכון המבני הגדול: **אין אימות בצד השרת** — `{userId, level}` מהלקוח. לכל endpoint חדש/ששונה: האם משתמש ב-gate של
  אחיו; האם בדיקת בעלות (`GetCompanyIdByLevelAndUserId`) משתמשת ב-id שה-query באמת עובד איתו; האם endpoint נהג מציג
  נתונים של נהג אחר.
- `back/.env` = production RDS; task-definition files — לא נקראים. CORS מתיר כל `http://localhost:<port>`.
- אל תציעו "תיקון" שמניח אימות שלא קיים; סמנו זאת כממצא מבני וציינו את `IDEAS.md` (אימות בצד השרת).
- ממצא חדש אמיתי → שורה מוצעת ל-§9 של ה-spec הרלוונטי (STATUS פתוח) ל-session הראשי; לא עורכים בעצמכם.
