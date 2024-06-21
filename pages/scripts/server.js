export const getData = async (url) => {
    const response = await fetch(url);

    if (!response.ok) {
        throw new Error(`Ошибка по адресу ${url}, статус ошибки ${response}`);
    }

    return await response.json();
};

export const sendDataPost = async (url, data) => {
    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)

        });

        if (!response.ok) {
            document.querySelector('.alert-container').innerHTML += 
            `
            <div class="alert alert-danger alert-dismissible fade show" style="display:inline-block;" role="alert">
                <strong>Ошибка!</strong>
                <p>${response.status}:${response.statusText}</p>
                <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
             </div>
            `
            throw new Error(`Ошибка по адресу ${url}, статус ошибки ${response.status}: ${response.statusText}`);

        }
        else{
            document.querySelector('.alert-container').innerHTML += 
            `
            <div  class="alert alert-success alert-dismissible fade show" style="display:inline-block;" role="alert">
                <strong>Успешно!</strong> 
                <p>${response.status}:${response.statusText}</p>
                <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
            </div>
            `
        }


        return await response;
    } catch (error) {
        console.error('Ошибка при отправке данных:', error);
        throw error;
    }
};

export const sendDataPut = async (url, data) => {
    try {
        const response = await fetch(url, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)

        });

        if (!response.ok) {
            document.querySelector('.alert-container').innerHTML += 
            `
            <div class="alert alert-danger alert-dismissible fade show" style="display:inline-block;" role="alert">
                <strong>Ошибка!</strong>
                <p>${response.status}:${response.statusText}</p>
                <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
             </div>
            `
            throw new Error(`Ошибка по адресу ${url}, статус ошибки ${response.status}: ${response.statusText}`);

        }
        else{
            document.querySelector('.alert-container').innerHTML += 
            `
            <div  class="alert alert-success alert-dismissible fade show" style="display:inline-block;" role="alert">
                <strong>Успешно!</strong> 
                <p>${response.status}:${response.statusText}</p>
                <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
            </div>
            `
        }


        return await response;
    } catch (error) {
        console.error('Ошибка при отправке данных:', error);
        throw error;
    }
};

export const sendDataDelete = async (url) => {
    try {
        const response = await fetch(url, {
            method: 'DELETE'
        });

        if (!response.ok) {
            document.querySelector('.alert-container').innerHTML += 
            `
            <div class="alert alert-danger alert-dismissible fade show" style="display:inline-block;" role="alert">
                <strong>Ошибка!</strong>
                <p>${response.status}:${response.statusText}</p>
                <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
             </div>
            `
            throw new Error(`Ошибка по адресу ${url}, статус ошибки ${response.status}: ${response.statusText}`);

        }
        else{
            document.querySelector('.alert-container').innerHTML += 
            `
            <div  class="alert alert-success alert-dismissible fade show" style="display:inline-block;" role="alert">
                <strong>Успешно!</strong> 
                <p>${response.status}:${response.statusText}</p>
                <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
            </div>
            `
        }


        return await response;
    } catch (error) {
        console.error('Ошибка при отправке данных:', error);
        throw error;
    }
};