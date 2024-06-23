import { getData } from './server.js'; 
import { sendDataPost } from './server.js';

const currentTeacher = JSON.parse(localStorage.getItem('currentTeacher'));

// Получаем список всех файлов
async function getAllTableFile() {
    const url = '/tableFile/all';
    const data = await getData(url);
    return data;
  }
  

async function fillSubjectSelect() {
    const tableFileSelect = document.getElementById('table-file-select');
    const allTableFile = await getAllTableFile();
    allTableFile.forEach(tableFile => {
      const option = document.createElement('option');
      option.value = tableFile.ID;
      option.textContent = tableFile.name;
      tableFileSelect.appendChild(option);
    });
  }

  document.getElementById('add-pricing-form').addEventListener('submit', async event => {
    event.preventDefault();
    const tableFileId = document.getElementById('table-file-select').value;

    const formDataObj = {
        ID: parseInt(tableFileId, 10)
      };
    
    sendDataPost("/pricing/add", formDataObj)
    .then(() => {
        window.location.href = 'get.html';
    })
    .catch((err) => {
        console.log(err);
    });


  });

fillSubjectSelect();