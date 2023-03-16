
const formElement = document.getElementById('property-inspector');

window.addEventListener('DOMContentLoaded', () => {
  window.form.forEach(field => {
    switch (field.type) {
      case 'text': {
        const fieldContainer = document.createElement('div');
        fieldContainer.classList.add('sdpi-item');

        const fieldLabel = document.createElement('div');
        fieldLabel.classList.add('sdpi-item-label');
        fieldLabel.title = field.label;
        fieldLabel.textContent = field.label;

        const fieldInput = document.createElement('input');
        fieldInput.classList.add('sdpi-item-value');
        fieldInput.type = 'text';
        fieldInput.name = field.name;
        fieldInput.placeholder = field.placeholder ?? '';
        fieldInput.value = field.defaultValue ?? '';

        fieldContainer.append(fieldLabel);
        fieldContainer.append(fieldInput);

        formElement.append(fieldContainer);

        break;
      }
      default:
        break;
    }
  });
});
