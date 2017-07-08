export function RegisterModel(model) {
  if (!(global.DvaApp._models.filter(m => m.namespace === model.namespace).length === 1)) {
    global.DvaApp.model(model);
  }
}
