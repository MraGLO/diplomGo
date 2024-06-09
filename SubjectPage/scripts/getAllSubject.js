import { getData } from './server.js'; 

export async function parseData(url, selectName) {
    const data = await getData(url);
    const selectElement = document.querySelector(`div[name="${selectName}"]`);

    data.forEach(item => {
        selectElement.insertAdjacentHTML('beforeend', `
        <div class="row">
                <div class="col-1 border">
                    ${item.id}
                </div>
                <div class="col-7 border">
                    ${item.subjectName}
                </div>
                    <div class="col-2 border" align="center">
                        <a type="button" class="btn btn-warning" style="width: 100px; margin: 5px;">Изменить</a>
                    </div>
                    <div class="col-2 border" align="center">
                        <a type="button" class="btn btn-danger" style="width: 100px; margin: 5px;">Удалить</a>
                    </div>
            </div>
        </div>
        `);
    });
}
parseData('/subject/all', 'allSubject');