package slack

type Channel struct {
	ID   string
	Name string
}

func FindChannelByID(channels []Channel, id string) Channel {
	for _, c := range channels {
		if c.ID == id {
			return c
		}
	}
	return nil
}
