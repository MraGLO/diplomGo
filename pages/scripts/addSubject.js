import { sendDataPost } from './server.js';

window.onload = function() {
    document.getElementById('btnE').addEventListener('click', function(event) {
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
      sendDataPost("/subject/add", formDataObj)
      .then(() => {
        window.location.href = 'get.html';
      })
      .catch((err) => {
          console.log(err);
      });
   });
}
