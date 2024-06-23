import { getData } from './server.js'; 
import { sendDataPost } from './server.js';

const currentPricing = JSON.parse(localStorage.getItem('currentPricing'));
document.getElementById('title').textContent = `Добавление преподавателя в таблицу ${currentPricing.name}`;


// Получаем список всех предметов
async function getAllTeachers() {
    const url = '/teacher/all';
    const data = await getData(url);
    return data;
  }
  
  // Получаем список предметов, которые уже преподает текущий учитель
  async function getPricingTeacher() {
    const id = currentPricing.id;
  const url = `/pricingTeacher/${id}`;
  try {
    const data = await getData(url);
    return data.map(item => item.ID);
  } catch (error) {
    console.error(error);
    return [];
  }
  }



async function fillSubjectSelect() {
    const teacherSelect = document.getElementById('teacher-select');
    const allTeachers = await getAllTeachers();
    const pricingTeacher = await getPricingTeacher();
    const availableSubjects = allTeachers.filter(teacher => !pricingTeacher.includes(teacher.ID));
    availableSubjects.forEach(teacher => {
      const option = document.createElement('option');
      option.value = teacher.ID;
      option.textContent = `${teacher.name} ${teacher.surname} ${teacher.patronymic}`;
      teacherSelect.appendChild(option);
    });
  }

  document.getElementById('teacher-pricing-form').addEventListener('submit', async event => {
    event.preventDefault();
    const teacherId = document.getElementById('teacher-select').value;

    const formDataObj = {
        teacherID: parseInt(teacherId, 10),
        pricingID: parseInt(currentPricing.id, 10) 
      };
    sendDataPost("/pricingTeacher/add", formDataObj)
    .then(() => {
        window.location.href = 'get.html';
    })
    .catch((err) => {
        console.log(err);
    });


  });

fillSubjectSelect();