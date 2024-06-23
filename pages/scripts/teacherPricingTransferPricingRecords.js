const table = document.getElementById('my-table');

// Add event listener to each "Get Value" button
table.addEventListener('click', (e) => {
  if (e.target.classList.contains('get-pricing-teacher-btn-tr')) {
    const row = e.target.parentElement.parentElement;
    const cells = row.cells;

    var pricingTeacher ={
        id: cells[0].textContent,
        fullname: cells[1].textContent
    }

    // Сохраняем текущего студента в localStorage
    localStorage.setItem('currentPricingTeacher', JSON.stringify(pricingTeacher));
    // Открываем страницу с предметами
    window.location.href = '/pricing_teacher_records/get.html';
  }
});
