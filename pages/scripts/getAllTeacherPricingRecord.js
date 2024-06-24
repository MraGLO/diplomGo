async function parseData(selectName, formDataObj) {
    try {
      const response = await fetch('/pricingTeacherRecord/all', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(formDataObj)
      });
  
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
  
      const data = await response.json();
  
      const selectElement = document.querySelector(`tbody[name="${selectName}"]`);
  
      data.forEach(item => {
          selectElement.insertAdjacentHTML('beforeend', `
              <tr>
                  <td>${item.ID}</td>
                  <td>${item.group}</td>
                  <td>${item.firstHalfYear}</td>
                  <td>${item.secondHalfYear}</td>
                  <td>${item.theory}</td>
                  <td>${item.practice}</td>
                  <td>${item.lpz1}</td>
                  <td>${item.lpz2}</td>
                  <td>${item.consultation}</td>
                  <td>${item.courseProject}</td>
                  <td>${item.hoursFirstSemester}</td>
                  <td>${item.hoursSecondSemester}</td>
                  <td>${item.total}</td>
                  <td>${item.subject}</td>
                <td>
                    <button type="button" class="btn btn-warning get-teacher-btn-u" style="width: 100px; margin: 5px;">Изменить</button>
                </td>
                <td>
                    <button type="button" class="btn btn-danger get-teacher-btn-d"  style="width: 100px; margin: 5px;">Удалить</button>
                </td>
              </tr>
          `);
      });
    } catch (error) {
      console.error('There was a problem with the fetch operation:', error);
    }
  }
  
  // Получаем текущего студента из localStorage
  const currentTeacher = JSON.parse(localStorage.getItem('currentPricingTeacher'));
  const currentPricing = JSON.parse(localStorage.getItem('currentPricing'));
  
  const formDataObj = {
      teacherID: parseInt(currentTeacher.id, 10),
      pricingID: parseInt(currentPricing.id, 10)
  };
  
  // Выводим имя и фамилию студента в заголовок
  document.getElementById('title').textContent = `${currentTeacher.fullname}`;
  
  parseData('allTeacherPricingRecord', formDataObj);