package model.core;

import model.utility.Title;

public class Snippet extends Media {
    Title title;
    Media parent;
    int starttime;
    int length;

    Snippet(Title title, Media parent, int start, int length) {
        super(title);
        this.title = title;
        this.parent = parent;
        this.starttime = start;
        this.length = length;
    }

    @Override
    public Snippet createSnippet(Title title, int start, int length) {
        return new Snippet(title, this.parent, this.starttime + start, length);
    }
}
