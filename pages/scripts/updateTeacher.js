import { sendDataPut } from './server.js';

window.onload = function() {
    document.getElementById('btnE').addEventListener('click', function(event) {
      event.preventDefault();
      // Функционал обработки клика

      const publicationForm = document.querySelector('.container');

      // Собираем данные из формы в объект
      

      // Сбор данных из формы
      const formData = new FormData(publicationForm);

      const select = document.getElementById('category-select');
      const selectedOptions = select.selectedOptions;
      var optionValue
    
      for (const option of selectedOptions) {
        optionValue = option.value;
      }

      const formDataObj = {
        ID: parseInt(formData.get('ID'), 10),
        name: formData.get('name'),
        surname: formData.get('surname'),
        patronymic: formData.get('patronymic'),
        category: parseInt(optionValue, 10)
      };

      const textField = document.getElementById('id-input');
      const value = textField.value;

      // Отправляем данные на сервер
      sendDataPut(`/teacher/update/${value}`, formDataObj)
      .then(() => {
          publicationForm.reset();
      })
      .catch((err) => {
          console.log(err);
      });
   });
}
