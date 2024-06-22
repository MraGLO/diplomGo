import { sendDataDelete } from './server.js';

window.onload = function() {
    document.getElementById('btnE').addEventListener('click', function(event) {
      event.preventDefault();
      // Функционал обработки клика

      const currentTeacher = JSON.parse(localStorage.getItem('currentTeacher'));
      const textField = document.getElementById('id-input');
      const id = textField.value;

      const formDataObj = {
        ID: null,
        teacherID: parseInt(currentTeacher.id, 10),
        subjectID: parseInt(id, 10) 
      };

      // Отправляем данные на сервер
      sendDataDelete(`/teacherSubject/delete`, formDataObj)
      .then(() => {
          publicationForm.reset();
      })
      .catch((err) => {
          console.log(err);
      });
   });
}
