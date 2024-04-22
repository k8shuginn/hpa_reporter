package collector

// OnAddFunc is hpa event on add function
type OnAddFunc func(interface{}, bool)

// OnAdd is hpa event on add
func (add OnAddFunc) OnAdd(obj interface{}, isInInitialList bool) {
	if add == nil {
		return
	}

	add(obj, isInInitialList)
}

// OnUpdateFunc is hpa event on update function
type OnUpdateFunc func(interface{}, interface{})

// OnUpdate is hpa event on update
func (update OnUpdateFunc) OnUpdate(oldObj, newObj interface{}) {
	if update == nil {
		return
	}

	update(oldObj, newObj)
}

// OnDeleteFunc is hpa event on delete function
type OnDeleteFunc func(interface{})

// OnDelete is hpa event on delete
func (delete OnDeleteFunc) OnDelete(obj interface{}) {
	if delete == nil {
		return
	}

	delete(obj)
}
