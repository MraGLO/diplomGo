const table = document.getElementById('my-table');

// Add event listener to each "Get Value" button
table.addEventListener('click', (e) => {
  if (e.target.classList.contains('get-teacher-btn-ts')) {
    const row = e.target.parentElement.parentElement;
    const cells = row.cells;

    var teacher ={
        id: cells[0].textContent,
        name: cells[1].textContent,
        surname: cells[2].textContent,
        patronymic: cells[3].textContent
    }
    const id = cells[0].textContent;

    // Сохраняем текущего студента в localStorage
    localStorage.setItem('currentTeacher', JSON.stringify(teacher));
    // Открываем страницу с предметами
    window.location.href = '/teacher_subject/get.html';
  }
});
