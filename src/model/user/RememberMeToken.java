package model.user;

public class RememberMeToken implements Credentials {
    String token;

    @Override
    public Session getSession() {
        return null;
    }

    @Override
    public boolean isValid() {
        return false;
    }
}
