const table = document.getElementById('my-table');

// Add event listener to each "Get Value" button
table.addEventListener('click', (e) => {
  if (e.target.classList.contains('get-teacher-btn-u')) {
    const row = e.target.parentElement.parentElement;
    const cells = row.cells;
    const categoryEnum = {
        1: 'Нет',
        2: 'Первая',
        3: 'Высшая'
      };
    // Get the values of the ID and Name columns
    const id = cells[0].textContent;
    const name = cells[2].textContent;
    const surname = cells[1].textContent;
    const patronymic = cells[3].textContent;
    const category = cells[4].textContent;

    // Create a new URL with the data as query parameters
    const url = `update.html?id=${encodeURIComponent(id)}&name=${encodeURIComponent(name)}&surname=${encodeURIComponent(surname)}&patronymic=${encodeURIComponent(patronymic)}&category=${Object.keys(categoryEnum).find(k => categoryEnum[k] === category)}`;

    // Redirect to the new page
    window.location.href = url;

  }
});
