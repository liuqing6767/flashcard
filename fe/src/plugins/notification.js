// https://2ue.github.io/2017/10/16/desktop-notification/

function notification(handler, interval) {
	if (!window.Notification) {
		return;
	}
	// "granted" "default" "denied"
	if (Notification.permission == "denied") {
		return;
	}
	if (Notification.permission == "granted") {
		if (interval) {
			setInterval(() => {
				handler();
			}, interval);
		} else {
			handler();
		}
		return;
	}

	Notification.requestPermission().then(permission => {
		if (permission == "granted") {
			notification(handler, interval);
		}
	});
}

export {
	notification
};