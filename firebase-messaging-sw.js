importScripts('https://www.gstatic.com/firebasejs/9.6.1/firebase-app-compat.js');
importScripts('https://www.gstatic.com/firebasejs/9.6.1/firebase-messaging-compat.js');

firebase.initializeApp({
    apiKey: "AIzaSyA55jHSEgbOycz9Y_vcrO0bx0wSw6Xx5wc",
    authDomain: "monitor-614c1.firebaseapp.com",
    projectId: "monitor-614c1",
    messagingSenderId: "1086285248664",
    appId: "1:1086285248664:web:1dff6c54eb37f193e238ed"
});

const messaging = firebase.messaging();

messaging.onBackgroundMessage(function(payload) {
    console.log('📥 Background message:', payload);
    const { title, body } = payload.notification;

    self.registration.showNotification(title, {
        body: body,
        icon: '/icon.png',
    });
});