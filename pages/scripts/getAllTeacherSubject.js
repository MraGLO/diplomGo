import { getData } from './server.js'; 

export async function parseData(url, selectName) {
    const data = await getData(url);

    const selectElement = document.querySelector(`tbody[name="${selectName}"]`);

    data.forEach(item => {
        selectElement.insertAdjacentHTML('beforeend', `
            <tr>
                <td>${item.ID}</td>
                <td>${item.subjectName}</td>
                <td>
                    <button type="button" class="btn btn-danger get-table-file-btn-d"  style="width: 100px; margin: 5px;">Удалить</button>
                </td>
            </tr>
        `);
    });
}

// Получаем текущего студента из localStorage
const currentTeacher = JSON.parse(localStorage.getItem('currentTeacher'));

// Выводим имя и фамилию студента в заголовок
document.getElementById('title').textContent = `${currentTeacher.surname} ${currentTeacher.name} ${currentTeacher.patronymic}`;

var url = `/teacher_subject/${currentTeacher.id}`

parseData(url, 'allTeacherSubject');