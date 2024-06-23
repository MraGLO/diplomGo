import { sendDataPost } from './server.js';

window.onload = function() {
    document.getElementById('btnE').addEventListener('click', function(event) {
      event.preventDefault();
      // Функционал обработки клика

      const publicationForm = document.querySelector('.container');

      // Собираем данные из формы в объект
      const formData = new FormData(publicationForm);

      const select = document.getElementById('category-select');
      const selectedOptions = select.selectedOptions;
      var optionValue
    
      for (const option of selectedOptions) {
        optionValue = option.value;
      }

      const formDataObj = {
        ID: parseInt(formData.get('ID'), 10),
        surname: formData.get('surname'),
        name: formData.get('name'),
        patronymic: formData.get('patronymic'),
        category: parseInt(optionValue, 10)
      };


      // Отправляем данные на сервер
      sendDataPost("/teacher/add", formDataObj)
      .then(() => {
            window.location.href = 'get.html';
      })
      .catch((err) => {
          console.log(err);
      });
   });

   
}
