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
                    <button type="button" class="btn btn-warning get-value-btn-u" href="update.html" style="width: 100px; margin: 5px;">Изменить</button>
                </td>
                <td>
                    <button type="button" class="btn btn-danger get-value-btn-d" href="delete.html" style="width: 100px; margin: 5px;">Удалить</button>
                </td>
            </tr>
        `);
    });
}
parseData('/subject/all', 'allSubject');