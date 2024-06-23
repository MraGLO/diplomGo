import { sendDataDelete } from './server.js';

const currentPricing = JSON.parse(localStorage.getItem('currentPricing'));
document.getElementById('title').textContent = `Удаление преподавателя из таблицы ${currentPricing.name}`;

const currentPricingTeacher = JSON.parse(localStorage.getItem('currentPricingTeacher'));
document.getElementById('id-input').value = currentPricingTeacher.id;
document.getElementById('FIO-input').value = currentPricingTeacher.fullname;


  document.getElementById('teacher-pricing-form').addEventListener('submit', async event => {
    event.preventDefault();

    const formDataObj = {
        teacherID: parseInt(currentPricingTeacher.id, 10),
        pricingID: parseInt(currentPricing.id, 10) 
      };
    sendDataDelete("/pricingTeacher/delete", formDataObj)
    .then(() => {
        window.location.href = 'get.html';
    })
    .catch((err) => {
        console.log(err);
    });

  });
