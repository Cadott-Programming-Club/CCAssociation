// Calendar interactions: HTMX-loaded day modal, view toggle, native share,
// copy link, add-to-calendar dropdown auto-close.
(function () {
  'use strict';

  // ── Day modal: HTMX swaps content into #cal-modal, then we open it.
  document.body.addEventListener('htmx:afterSwap', function (e) {
    if (e.target && e.target.id === 'cal-modal') {
      var dlg = document.getElementById('cal-modal');
      if (dlg && typeof dlg.showModal === 'function' && !dlg.open) {
        dlg.showModal();
      }
    }
  });

  // Backdrop click closes <dialog>.
  document.addEventListener('click', function (e) {
    var dlg = document.getElementById('cal-modal');
    if (!dlg || !dlg.open || e.target !== dlg) return;
    var r = dlg.getBoundingClientRect();
    var outside =
      e.clientX < r.left ||
      e.clientX > r.right ||
      e.clientY < r.top ||
      e.clientY > r.bottom;
    if (outside) dlg.close();
  });

  // ── View toggle: persists in localStorage, updates URL via history.
  var toggle = document.querySelector('[data-view-toggle]');
  if (toggle) {
    toggle.addEventListener('click', function (e) {
      var btn = e.target.closest('[data-view]');
      if (!btn) return;
      var view = btn.getAttribute('data-view');
      try { localStorage.setItem('calendarView', view); } catch (_) {}
      var url = new URL(window.location.href);
      url.searchParams.set('view', view);
      window.location.href = url.toString();
    });
  }

  // On first paint of /events/calendar, jump to the saved view if no
  // explicit ?view= is set.
  if (location.pathname === '/events/calendar' && !new URL(location.href).searchParams.has('view')) {
    try {
      var saved = localStorage.getItem('calendarView');
      if (saved === 'list' || saved === 'month') {
        var u = new URL(location.href);
        u.searchParams.set('view', saved);
        // Replace (not push) so back button skips the redirect.
        location.replace(u.toString());
      }
    } catch (_) {}
  }

  // ── Share buttons: native share where available; otherwise fall back
  // to the Facebook/Email/Copy row already in the DOM.
  document.addEventListener('click', function (e) {
    var nativeBtn = e.target.closest('[data-share-native]');
    if (nativeBtn) {
      var root = nativeBtn.closest('[data-share]');
      if (!root) return;
      var data = {
        title: root.getAttribute('data-share-title') || '',
        text: root.getAttribute('data-share-text') || '',
        url: root.getAttribute('data-share-url') || location.href,
      };
      if (navigator.share) {
        e.preventDefault();
        navigator.share(data).catch(function () {/* user dismissed */});
        return;
      }
      // No native share — let the fallback row do the work; nothing to do.
      return;
    }

    var copyBtn = e.target.closest('[data-share-copy]');
    if (copyBtn) {
      e.preventDefault();
      var root2 = copyBtn.closest('[data-share]');
      var url = (root2 && root2.getAttribute('data-share-url')) || location.href;
      if (navigator.clipboard && navigator.clipboard.writeText) {
        navigator.clipboard.writeText(url).then(function () {
          flashToast('Link copied');
        });
      } else {
        var ta = document.createElement('textarea');
        ta.value = url;
        document.body.appendChild(ta);
        ta.select();
        try { document.execCommand('copy'); flashToast('Link copied'); } catch (_) {}
        ta.remove();
      }
    }
  });

  function flashToast(msg) {
    var t = document.createElement('div');
    t.textContent = msg;
    t.setAttribute('role', 'status');
    t.style.cssText =
      'position:fixed;left:50%;bottom:1.5rem;transform:translateX(-50%);' +
      'background:#1A2332;color:#fff;padding:.5rem 1rem;border-radius:.5rem;' +
      'font-size:.875rem;z-index:1000;box-shadow:0 4px 12px rgba(0,0,0,.2);';
    document.body.appendChild(t);
    setTimeout(function () { t.remove(); }, 1800);
  }

  // ── Add-to-calendar <details>: close when clicking outside.
  document.addEventListener('click', function (e) {
    document.querySelectorAll('[data-add-to-calendar][open]').forEach(function (d) {
      if (!d.contains(e.target)) d.removeAttribute('open');
    });
  });
})();
