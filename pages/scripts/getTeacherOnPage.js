const queryString = window.location.search;
const urlParams = new URLSearchParams(queryString);

const id = urlParams.get('id');
const name = urlParams.get('name');
const surname = urlParams.get('surname');
const patronymic = urlParams.get('patronymic');
const category = urlParams.get('category');

// Set the values of the input fields
document.getElementById('id-input').value = id;
document.getElementById('name-input').value = name;
document.getElementById('surname-input').value = surname;
document.getElementById('patronymic-input').value = patronymic;

const select = document.getElementById('category-select');
const options = select.querySelectorAll(`option[value="${category}"]`);

options.forEach(option => {
  option.selected = true;
});
