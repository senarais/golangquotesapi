package database

import (
	"github.com/senarais/golangquotesapi/internal/model"
)

type QuotesData struct {}

func(q *QuotesData) GetData() []model.Quotes {
	data := []model.Quotes{
		{Quote: "Stay hungry, stay foolish.", Author: "Steve Jobs"},
		{Quote: "The best way to get started is to quit talking and begin doing.", Author: "Walt Disney"},
		{Quote: "Don’t let yesterday take up too much of today.", Author: "Will Rogers"},
		{Quote: "It’s not whether you get knocked down, it’s whether you get up.", Author: "Vince Lombardi"},
		{Quote: "If you are working on something exciting, it will keep you motivated.", Author: "Steve Jobs"},
		{Quote: "Success is not in what you have, but who you are.", Author: "Bo Bennett"},
		{Quote: "Hard work beats talent when talent doesn’t work hard.", Author: "Tim Notke"},
		{Quote: "Dream bigger. Do bigger.", Author: "Unknown"},
		{Quote: "Do what you can, with what you have, where you are.", Author: "Theodore Roosevelt"},
		{Quote: "Everything you’ve ever wanted is on the other side of fear.", Author: "George Addair"},
		{Quote: "The harder you work for something, the greater you’ll feel when you achieve it.", Author: "Unknown"},
		{Quote: "Don’t watch the clock; do what it does. Keep going.", Author: "Sam Levenson"},
		{Quote: "Great things never come from comfort zones.", Author: "Unknown"},
		{Quote: "Push yourself, because no one else is going to do it for you.", Author: "Unknown"},
		{Quote: "Success doesn’t just find you. You have to go out and get it.", Author: "Unknown"},
		{Quote: "Dream it. Wish it. Do it.", Author: "Unknown"},
		{Quote: "Stay positive, work hard, make it happen.", Author: "Unknown"},
		{Quote: "Failure is not the opposite of success; it’s part of success.", Author: "Arianna Huffington"},
		{Quote: "Don’t stop when you’re tired. Stop when you’re done.", Author: "Marilyn Monroe"},
		{Quote: "Wake up with determination. Go to bed with satisfaction.", Author: "Unknown"},
	}

	return data
}