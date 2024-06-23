const publicationForm = document.querySelector('.container');
publicationForm.addEventListener('submit', e => {
        e.preventDefault();

        let excel = document.getElementById("customFile").files[0];
        let formData = new FormData();
                
        formData.append("excel", excel);

        let r = fetch('/tableFile/add', {method: "POST", body: formData}); 
        console.log('HTTP response code:',r.status);
});
