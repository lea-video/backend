package model.core;

import model.utility.Tag;
import model.utility.Title;

import java.util.ArrayList;
import java.util.List;

public class Playlist implements Playable {
    List<Tag> tags;
    Title title;
    List<PlaylistItem> items;

    public Playlist(Title title) {
        this.title = title;
        this.tags = new ArrayList<>();
        this.items = new ArrayList<>();
    }

    public PlaylistItem addItem(Playable m) {
        PlaylistItem pi = new PlaylistItem(this.items.size(), m);
        this.items.add(pi);
        return pi;
    }

    @Override
    public Title getTitle() {
        return this.title;
    }

    @Override
    public Playable getNext(int current) {
        if (current + 1 >= this.items.size()) {
            return null;
        }
        return items.get(current + 1).playable;
    }
}
