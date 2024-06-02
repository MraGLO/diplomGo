import { sendData } from './server.js';

// отправка на сервер
const publicate = (url) => {
    const publicationForm = document.querySelector('.publication-form');

    publicationForm.addEventListener('submit', e => {
        e.preventDefault();

       // Собираем данные из формы в объект
       const formDataObj = {};

       // Сбор данных из формы
       const formData = new FormData(publicationForm);
       for (let [key, value] of formData.entries()) {
           formDataObj[key] = value;
       }

       // Отправляем данные на сервер
       sendData(url, formDataObj)
       .then(() => {
           publicationForm.reset();
       })
       .catch((err) => {
           console.log(err);
       });
    });
}

publicate('/file/add');