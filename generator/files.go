package generator

func (g *Generator) ClientFile() (Templater, error) {
	// var client ClientOld
	// for _, o := range g.Operations {
	// 	client.ClientHandlers = append(client.ClientHandlers, NewClientHandlerOld(o))
	// }
	return g.goFile([]string{}, g.Client), nil
}
