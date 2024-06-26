const table = document.getElementById('my-table');

// Add event listener to each "Get Value" button
table.addEventListener('click', (e) => {
  if (e.target.classList.contains('get-subject-btn-u')) {
    const row = e.target.parentElement.parentElement;
    const cells = row.cells;

    // Get the values of the ID and Name columns
    const id = cells[0].textContent;
    const name = cells[1].textContent;

    // Create a new URL with the data as query parameters
    const url = `update.html?id=${encodeURIComponent(id)}&name=${encodeURIComponent(name)}`;

    // Redirect to the new page
    window.location.href = url;

  }
});
