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
                    <button type="button" class="btn btn-info get-pricing-table-btn-tt" style="width: 100px; margin: 5px;">Данные</a>
                </td>
                <td>
                    <button type="button" class="btn btn-warning get-pricing-table-btn-u" style="width: 100px; margin: 5px;">Изменить</button>
                </td>
                <td>
                    <button type="button" class="btn btn-danger get-pricing-table-btn-d" style="width: 100px; margin: 5px;">Удалить</button>
                </td>
                <td>
                    <button type="button" class="btn btn-success " style="width: 100px; margin: 5px;">Экспорт</button>
                </td>
            </tr>
        `);
    });
}
parseData('/pricing/all', 'allPricing');