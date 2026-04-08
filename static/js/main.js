// Calculate the Friday of the last full weekend in July for the given year.
// "Last full weekend" = last Sat+Sun pair that both fall in July.
// Nabor Days starts on the Friday before that Saturday.
function getNaborDaysFriday(year) {
  // Find the last day of July
  // Walk backwards from July 31 to find the last Saturday
  let d = new Date(year, 6, 31); // July 31 (month is 0-indexed)
  while (d.getDay() !== 6) { // 6 = Saturday
    d.setDate(d.getDate() - 1);
  }
  // d is now the last Saturday in July
  // Check if Sunday (d+1) is also in July
  let sunday = new Date(d);
  sunday.setDate(sunday.getDate() + 1);
  if (sunday.getMonth() !== 6) {
    // Sunday falls in August, so go back one week
    d.setDate(d.getDate() - 7);
  }
  // d is now the Saturday of the last full weekend in July
  // Nabor Days starts Friday (one day before)
  let friday = new Date(d);
  friday.setDate(friday.getDate() - 1);
  // Set to 5 PM (festival start)
  friday.setHours(17, 0, 0, 0);
  return friday;
}

function getNextNaborDays() {
  const now = new Date();
  let target = getNaborDaysFriday(now.getFullYear());
  // If this year's Nabor Days has passed, use next year
  if (now > target) {
    target = getNaborDaysFriday(now.getFullYear() + 1);
  }
  return target;
}

// Countdown timer
function updateCountdown() {
  const el = document.getElementById('countdown');
  if (!el) return;

  const target = getNextNaborDays().getTime();
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

// Mobile nav
function setupMobileNav() {
  const toggle = document.getElementById('nav-toggle');
  const menu = document.getElementById('nav-menu');
  if (!toggle || !menu) return;

  menu.querySelectorAll('a').forEach(link => {
    link.addEventListener('click', () => {
      menu.classList.add('hidden');
      menu.classList.remove('flex');
    });
  });

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
