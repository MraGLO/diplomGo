import { getData } from './server.js'; 

export async function parseData(url, selectName) {
    const data = await getData(url);

    const selectElement = document.querySelector(`tbody[name="${selectName}"]`);

    data.forEach(item => {
        selectElement.insertAdjacentHTML('beforeend', `
            <tr>
                <td>${item.ID}</td>
                <td>${item.name} ${item.surname} ${item.patronymic}</td>
                <td>
                    <button type="button" class="btn btn-info get-pricing-teacher-btn-tr"  style="width: 100px; margin: 5px;">Данные</button>
                </td>
                <td>
                    <button type="button" class="btn btn-danger get-pricing-teacher-btn-d"  style="width: 100px; margin: 5px;">Удалить</button>
                </td>
            </tr>
        `);
    });
}

// Получаем текущего студента из localStorage
const currentPricing = JSON.parse(localStorage.getItem('currentPricing'));

// Выводим имя и фамилию студента в заголовок
document.getElementById('title').textContent = `${currentPricing.name}`

var url = `/pricingTeacher/${currentPricing.id}`

parseData(url, 'allPricingTeacher');