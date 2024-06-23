const table = document.getElementById('my-table');

// Add event listener to each "Get Value" button
table.addEventListener('click', (e) => {
  if (e.target.classList.contains('get-pricing-table-btn-tt')) {
    const row = e.target.parentElement.parentElement;
    const cells = row.cells;

    var pricing ={
        id: cells[0].textContent,
        name: cells[1].textContent
    }

    // Сохраняем текущего студента в localStorage
    localStorage.setItem('currentPricing', JSON.stringify(pricing));
    // Открываем страницу с предметами
    window.location.href = '/pricing_teacher/get.html';
  }
});
