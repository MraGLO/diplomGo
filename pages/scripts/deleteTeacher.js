import { sendDataDelete } from './server.js';

window.onload = function() {
    document.getElementById('btnE').addEventListener('click', function(event) {
      event.preventDefault();
      // Функционал обработки клика


      const textField = document.getElementById('id-input');
      const value = textField.value;

      // Отправляем данные на сервер
      sendDataDelete(`/teacher/delete/${value}`)
      .then(() => {
          publicationForm.reset();
      })
      .catch((err) => {
          console.log(err);
      });
   });
}
