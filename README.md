# api-students
API to manage 'Golang do Zero' course students

Routes:
- GET /students - List all students
- GET /students/:id - Get infos from a specific student
- POST /students/:id - Update student
- DELETE /students/:id - Delete student

Struct Student:
- Name
- CPF
- Email
- Age
- Active