import { sendIDDelete } from './server.js';

window.onload = function() {
    document.getElementById('btnE').addEventListener('click', function(event) {
      event.preventDefault();
      // Функционал обработки клика


      const textField = document.getElementById('id-input');
      const value = textField.value;

      // Отправляем данные на сервер
      sendIDDelete(`/pricing/delete/${value}`)
      .then(() => {
        window.location.href = 'get.html';
      })
      .catch((err) => {
          console.log(err);
      });
   });
}
