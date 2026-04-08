// Countdown timer
function updateCountdown() {
  const el = document.getElementById('countdown');
  if (!el) return;

  const target = new Date(el.dataset.target).getTime();
  const now = Date.now();
  const diff = target - now;

  if (diff <= 0) {
    document.getElementById('countdown-days').textContent = '0';
    document.getElementById('countdown-hours').textContent = '0';
    document.getElementById('countdown-mins').textContent = '0';
    return;
  }

  const days = Math.floor(diff / (1000 * 60 * 60 * 24));
  const hours = Math.floor((diff % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
  const mins = Math.floor((diff % (1000 * 60 * 60)) / (1000 * 60));

  document.getElementById('countdown-days').textContent = days;
  document.getElementById('countdown-hours').textContent = hours;
  document.getElementById('countdown-mins').textContent = mins;
}

// Footer year
function setFooterYear() {
  const el = document.getElementById('footer-year');
  if (el) el.textContent = new Date().getFullYear();
}

// Mobile nav - close on route change
function setupMobileNav() {
  const toggle = document.getElementById('nav-toggle');
  const menu = document.getElementById('nav-menu');
  if (!toggle || !menu) return;

  // Close menu when clicking a nav link
  menu.querySelectorAll('a').forEach(link => {
    link.addEventListener('click', () => {
      menu.classList.add('hidden');
      menu.classList.remove('flex');
    });
  });

  // Close menu on resize to desktop
  window.addEventListener('resize', () => {
    if (window.innerWidth >= 640) {
      menu.classList.remove('hidden');
      menu.classList.add('flex');
    }
  });
}

// Init
document.addEventListener('DOMContentLoaded', () => {
  updateCountdown();
  setInterval(updateCountdown, 60000);
  setFooterYear();
  setupMobileNav();
});
