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
            throw new Error(`Ошибка по адресу ${url}, статус ошибки ${response.status}: ${response.statusText}`);

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
            throw new Error(`Ошибка по адресу ${url}, статус ошибки ${response.status}: ${response.statusText}`);
        }

        return await response;
    } catch (error) {
        console.error('Ошибка при отправке данных:', error);
        throw error;
    }
};

export const sendIDDelete = async (url) => {
    try {
        const response = await fetch(url, {
            method: 'DELETE'
        });

        if (!response.ok) {
            throw new Error(`Ошибка по адресу ${url}, статус ошибки ${response.status}: ${response.statusText}`);

        }

        return await response;
    } catch (error) {
        console.error('Ошибка при отправке данных:', error);
        throw error;
    }
};
export const sendDataDelete = async (url, data) => {
    try {
        const response = await fetch(url, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        if (!response.ok) {
            throw new Error(`Ошибка по адресу ${url}, статус ошибки ${response.status}: ${response.statusText}`);
        }


        return await response;
    } catch (error) {
        console.error('Ошибка при отправке данных:', error);
        throw error;
    }
};