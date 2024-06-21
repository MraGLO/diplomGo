const queryString = window.location.search;
const urlParams = new URLSearchParams(queryString);

const id = urlParams.get('id');
const name = urlParams.get('name');

// Set the values of the input fields
document.getElementById('id-input').value = id;
document.getElementById('name-input').value = name;