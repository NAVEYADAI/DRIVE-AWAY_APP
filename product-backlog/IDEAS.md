# רעיונות גולמיים (טרם אופיינו)

רעיונות לפיצ'רים, בערך שורה אחת לכל רעיון, **לפני** אפיון מלא. זו לא תחליף לקבצי הפיצ'רים
(`NN-slug.md`), אלא תחנת ביניים: כשמגיע הזמן לאפיין רעיון מכאן, מריצים עליו `product-manager` (עם
הרעיון בפרומפט). הוא הופך לקובץ פיצ'ר מלא עם מספר משלו, ואז השורה כאן מסומנת `[x]` עם קישור לקובץ.

אם רעיון עלול לגעת בפיצ'ר קיים (אותה טבלה/endpoint/קובץ), כדאי להריץ גם `feature-dependency-checker`.

אין כאן שום דבר ממומש ואין כאן החלטות מוצריות. הפריטים הראשונים הם **מסכים/endpoints שקיימים בקוד
כשלד בלבד** (נמצאו בזמן כתיבת `specs/`, 2026-09-29). העובדות מתועדות, ומה המוצר צריך לעשות בהם עוד
לא הוחלט.

---

- [ ] **הרשמה עצמית**: `font/src/LogIn/SingUp.tsx` ו-`MainLogIn.tsx` קיימים אבל אף route לא מציג אותם,
      ולכפתור השליחה אין `onClick`. ה-backend (`POST /loginApi/SingUp`) עובד, אבל לא יוצר שורת
      `permissions`, כך שמשתמש חדש לא יכול להתחבר דרך `POST /login`. טרם הוגדר: האם בכלל רוצים הרשמה
      עצמית (או שרק admin/manager יוצרים משתמשים), ואיזה level מקבל נרשם. ר' `specs/01-auth.md` §3.4, §6.

- [ ] **רישום עמדה בחברה קיימת (AddPosition) מקצה לקצה**: המסך
      `font/src/component/systemAdministrator/addPosition.tsx` מציג רשימת חברות **קשיחה**
      (`nave`, `yadai`), וכפתור השליחה לא קורא ל-API. ה-endpoint `POST /MainController/AddPosition`
      קיים. חסר endpoint לרשימת חברות (יש רק CRUD גולמי לא מוגן). ר' `specs/02-management.md` §3.1, §6.

- [ ] **דוח מספר שיבוצים לכל חברה (NumberOfInlays)**: המסך מציג נתוני דמו קשיחים, ואין endpoint
      (`systemadmin/NumberOfInlays.go` ריק). ר' `specs/02-management.md` §6.

- [ ] **עריכת שיבוצים (EditSort)**: מסך placeholder ("מסך זה נמצא בבנייה"). טרם הוגדר: מה מותר לערוך
      בשיבוץ קיים (להחליף נהג לנסיעה? למחוק?), והאם עריכה יוצרת `sort` חדש.

- [ ] **שחזור שם משתמש / שינוי שם משתמש וסיסמה**: `ForgetMyUserName` הוא stub שמחזיר 200 ריק.
      `ChangeUsername`/`ChangePassword` פונים ל-`localhost:8080` ולכן תמיד נכשלים. אין להם UI.
      ר' `specs/01-auth.md` §3.5–3.7, §9.4.

- [ ] **אימות אמיתי בצד השרת**: היום ה-backend סומך על `userId`/`level` שהלקוח שולח, ה-CRUD הגולמי
      לא מוגן, והסיסמאות נשמרות כטקסט גלוי. זו החלטה ארכיטקטונית/מוצרית (JWT? session? hashing?
      מה עושים עם ה-routes הגולמיים?) ולא תיקון נקודתי. ר' `specs/00-shared-conventions.md` §3.3, §9.

- [ ] **הפעלת E2E של יצירת שיבוץ**: `e2e/tests/sort-creation.spec.ts` מסומן `test.fixme` כי המיון קורא
      ל-Google Maps האמיתי. טרם הוגדר: stub מוזרק ל-`mapsHTTPClient` בסביבת E2E, או מפתח בדיקה.
      ר' `specs/03-sorting.md` §7.
