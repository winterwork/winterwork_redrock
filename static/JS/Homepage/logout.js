const homepage = document.getElementById('homepage');
const logout = document.getElementById('logout');
logout.addEventListener('click', async () => {
    window.localStorage.clear('token')
    alert('你已退出登录')
    location.reload();
})