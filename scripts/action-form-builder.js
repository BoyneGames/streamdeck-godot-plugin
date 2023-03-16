const fs = require('fs');
const actionForms = require('../action-forms.json');

const template = fs.readFileSync('./templates/property-inspector.html', 'utf8');

for (const action in actionForms) {
  const html = template.replace('%FORM%', JSON.stringify(actionForms[action]));
  fs.writeFileSync(`./games.boyne.godot.sdPlugin/html/${action}.html`, html);
}
