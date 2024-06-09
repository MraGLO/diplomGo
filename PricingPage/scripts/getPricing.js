import { getData } from './server.js'; 

export async function parseData(url, selectName) {
    const data = await getData(url);
    const selectElement = document.querySelector(`div[name="${selectName}"]`);

    data.forEach(item => {
        selectElement.insertAdjacentHTML('beforeend', `
        <div value="${item.pricingtable.id}">${item.pricingtable.name}</div>
        `);
    });
}
parseData('createPricing/5', 'pricing-form');