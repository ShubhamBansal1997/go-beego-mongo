/*
* @Author: Shubham Bansal
* @Date:   2018-10-02 15:33:29
* @Last Modified by:   Shubham Bansal
* @Last Modified time: 2018-10-02 17:57:18
*/
// package name
package models

// import go lib
import (
	"go-mongo/models/db"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Post model
type Post struct {
	Title string `json:"title" bson:"title"`
	SlugUrl string `json:"slug_url" bson:"slug_url"`
	Content string `json:"content" bson:"content"`
	PublishedAt time.Time `json:"published_at" bson:"published_at"`
	ID bson.ObjectId `json:"id" bson:"_id"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

func newPostCollection() *db.Collection {
	return db.NewCollectionSession("posts")
}

// create post method
func CreatePost(post Post) (Post, error) {
	var (
		err error
	)

	// Get post collection connection
	c := newPostCollection()
	defer c.Close()

	// set default mongodb ID and created date
	post.ID = bson.NewObjectId()
	post.CreatedAt = time.Now()

	// Insert post to mongodb
	err = c.Session.Insert(&post)
	if err != nil {
		return post, err
	}

	return post, err
}

// update post method
func (post Post) UpdatePost(postParam Post) (Post, error) {
	var (
		err error
	)
	// Get post collection connection
	c := newPostCollection()
	defer c.Close()

	// update post
	err = c.Session.Update(bson.M{
		"_id": post.ID,
		}, bson.M{
			"$set": bson.M{
				"title": postParam.Title,
				"slug_url": postParam.SlugUrl,
				"content": postParam.Content,
				"updated_at": time.Now(),
			},
		})
	if err != nil {
		return post, err
	}
	return post, err
}

// find posts method
func FindPosts() ( []Post, error) {
	var (
		err error
		posts []Post
	)

	// Get post collection connection
	c := newPostCollection()
	defer c.Close()

	// get posts
	err = c.Session.Find(nil).Sort("-published_at").All(&posts)
	if err != nil {
		return posts, err
	}
	return posts, err
}

// get post by id
func FindPost (id bson.ObjectId) (Post, error) {
	var (
		err error
		post Post
	)
	// Get post collection connection
	c := newPostCollection()
	defer c.Close()

	// get post
	err = c.Session.FindId(id).One(&post)
	if err != nil {
		return post, err
	}
	return post, err
}

// delete post by id
func DeletePost(post Post) (error) {
	var (
		err error
	)
	// Get post collection connection
	c := newPostCollection()
	defer c.Close()

	// remove Post
	err = c.Session.Remove(bson.M{"_id": post.ID})
	if err != nil {
		return err
	}
	return err
}

