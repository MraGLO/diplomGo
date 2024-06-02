const publicationForm = document.querySelector('.publication-form');
publicationForm.addEventListener('submit', e => {
        e.preventDefault();

        let excel = document.getElementById("file").files[0];
        let formData = new FormData();
                
        formData.append("excel", excel);

        let r = fetch('/file/add', {method: "POST", body: formData}); 
        console.log('HTTP response code:',r.status);
});
