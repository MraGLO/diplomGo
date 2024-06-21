import { sendDataPut } from './server.js';

window.onload = function() {
    document.getElementById('btnE').addEventListener('click', function(event) {
      event.preventDefault();
      // Функционал обработки клика

      const publicationForm = document.querySelector('.container');

      // Собираем данные из формы в объект
      

      // Сбор данных из формы
      const formData = new FormData(publicationForm);
      const formDataObj = {
        ID: parseInt(formData.get('ID'), 10),
        subjectName: formData.get('subjectName')
      };

      const textField = document.getElementById('id-input');
      const value = textField.value;

      // Отправляем данные на сервер
      sendDataPut(`/subject/update/${value}`, formDataObj)
      .then(() => {
          publicationForm.reset();
      })
      .catch((err) => {
          console.log(err);
      });
   });
}
