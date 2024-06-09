import { sendData } from './server.js';

window.onload = function() {
    document.getElementById('btnAdd').addEventListener('click', function(event) {
      event.preventDefault();
      // Функционал обработки клика

      const publicationForm = document.querySelector('.container');

      // Собираем данные из формы в объект
      const formDataObj = {};

      // Сбор данных из формы
      const formData = new FormData(publicationForm);
      for (let [key, value] of formData.entries()) {
          formDataObj[key] = value;
      }

      // Отправляем данные на сервер
      sendData("/subject/add", formDataObj)
      .then(() => {
          publicationForm.reset();
      })
      .catch((err) => {
          console.log(err);
      });
   });
}
