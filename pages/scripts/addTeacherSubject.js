import { getData } from './server.js'; 
import { sendDataPost } from './server.js';

const currentTeacher = JSON.parse(localStorage.getItem('currentTeacher'));

// Получаем список всех предметов
async function getAllSubjects() {
    const url = '/subject/all';
    const data = await getData(url);
    return data;
  }
  
  // Получаем список предметов, которые уже преподает текущий учитель
  async function getTeacherSubjects() {
    const id = currentTeacher.id;
  const url = `/teacher_subject/${id}`;
  try {
    const data = await getData(url);
    return data.map(item => item.subjectName);
  } catch (error) {
    console.error(error);
    return [];
  }
  }



async function fillSubjectSelect() {
    const subjectSelect = document.getElementById('subject-select');
    const allSubjects = await getAllSubjects();
    const teacherSubjects = await getTeacherSubjects();
    const availableSubjects = allSubjects.filter(subject => !teacherSubjects.includes(subject.subjectName));
    availableSubjects.forEach(subject => {
      const option = document.createElement('option');
      option.value = subject.ID;
      option.textContent = subject.subjectName;
      subjectSelect.appendChild(option);
    });
  }

  document.getElementById('add-subject-form').addEventListener('submit', async event => {
    event.preventDefault();
    const subjectId = document.getElementById('subject-select').value;

    const formDataObj = {
        ID: null,
        teacherID: parseInt(currentTeacher.id, 10),
        subjectID: parseInt(subjectId, 10) 
      };
    sendDataPost("/teacherSubject/add", formDataObj)
    .then(() => {
        publicationForm.reset();
    })
    .catch((err) => {
        console.log(err);
    });


  });

fillSubjectSelect();