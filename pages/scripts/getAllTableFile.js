import { getData } from './server.js'; 

export async function parseData(url, selectName) {
    const data = await getData(url);
    const selectElement = document.querySelector(`tbody[name="${selectName}"]`);
    
    data.forEach(item => {
        var date = new Date(item.date)
        var year = date.getFullYear();
        var month = (date.getMonth() + 1).toString().padStart(2, '0');
        var day = date.getDate().toString().padStart(2, '0');

        var formattedDate = `${year}-${month}-${day}`;
        selectElement.insertAdjacentHTML('beforeend', `
            <tr>
                <td>${item.ID}</td>
                <td>${item.name}</td>
                <td>${formattedDate}</td>
                <td>
                    <button type="button" class="btn btn-danger get-table-file-btn-d" style="width: 100px; margin: 5px;">Удалить</button>
                </td>
            </tr>
        `);
    });
}
parseData('/tableFile/all', 'allTableFile');