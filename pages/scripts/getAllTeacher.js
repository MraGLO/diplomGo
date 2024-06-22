import { getData } from './server.js'; 

export async function parseData(url, selectName) {
    const data = await getData(url);
    const selectElement = document.querySelector(`tbody[name="${selectName}"]`);
    const categoryEnum = {
        1: 'Нет',
        2: 'Первая',
        3: 'Высшая'
      };

    data.forEach(item => {
        selectElement.insertAdjacentHTML('beforeend', `
            <tr>
                <td>${item.ID}</td>
                <td>${item.surname}</td>
                <td>${item.name}</td>
                <td>${item.patronymic}</td>
                <td>${categoryEnum[item.category]}</td>
                <td>
                    <button type="button" class="btn btn-info" style="width: 100px; margin: 5px;">Предметы</a>
                </td>
                <td>
                    <button type="button" class="btn btn-warning get-teacher-btn-u" style="width: 100px; margin: 5px;">Изменить</button>
                </td>
                <td>
                    <button type="button" class="btn btn-danger get-teacher-btn-d"  style="width: 100px; margin: 5px;">Удалить</button>
                </td>
            </tr>
        `);
    });
}
parseData('/teacher/all', 'allTeacher');