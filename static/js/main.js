document.addEventListener("DOMContentLoaded", function() {
    const newsLinks = document.querySelectorAll('.news-link');

    newsLinks.forEach(link => {
        link.addEventListener('click', function(event) {
            event.preventDefault();
            const url = this.getAttribute('href');

            fetch(url)
                .then(response => response.text())
                .then(data => {
                    const detailContainer = document.getElementById('news-detail');
                    detailContainer.innerHTML = data;
                    detailContainer.scrollIntoView({ behavior: 'smooth' });
                })
                .catch(error => console.error('Error fetching news detail:', error));
        });
    });
});