document.addEventListener('DOMContentLoaded', () => {
    const deleteButtons = document.querySelectorAll('.delete-btn');

    deleteButtons.forEach(button => {
        button.addEventListener('click', async () => {
            const authorId = button.getAttribute('data-author-id');
            try {
                // Отправляем DELETE-запрос на сервер
                const response = await fetch(`/authors/${authorId}`, {
                    method: 'DELETE'
                });
                if (response.ok) {
                    window.location.reload();
                } else {
                    window.location.reload();
                }
            } catch (error) {
                console.error('Ошибка удаления автора:', error.message);
            }
        });
    });
});
