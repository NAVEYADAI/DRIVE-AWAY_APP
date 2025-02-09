# syntax=docker/dockerfile:1

# השתמש בתמונה הרשמית של Golang
FROM golang:1.21.0

# קבע את תיקיית העבודה בתוך הקונטיינר
WORKDIR /app

# העתק את קבצי go.mod ו-go.sum כדי להוריד את המודולים
COPY go.mod go.sum ./
RUN go mod download

# העתק את כל הקבצים כולל תיקיות משנה
COPY . ./

# הגדר את סביבת העבודה עבור הבנייה
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# הפעל את הפורט הנדרש
EXPOSE 2909

# הפעל את הקובץ המוכן
CMD ["/docker-gs-ping"]
