$PI.onConnected((jsn) => {
  const form = document.querySelector('#property-inspector');

  const { actionInfo } = jsn;
  const { payload } = actionInfo;
  const { settings } = payload;

  Utils.setFormValue(settings, form);

  form.addEventListener('input', Utils.debounce(150, () => {
    $PI.setSettings(Utils.getFormValue(form));
  }));
});
