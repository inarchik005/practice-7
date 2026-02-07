public class Time {

    private int hour;
    private int minute;
    private int second;

    public Time(int h, int m, int s) {
        hour = h;
        minute = m;
        second = s;
    }

    public String toUniversal() {
        return String.format("%02d:%02d:%02d", hour, minute, second);
    }

    public String toStandard() {
        int h = hour % 12;
        if (h == 0) h = 12;

        String period = hour < 12 ? "AM" : "PM";

        return String.format("%02d:%02d:%02d %s", h, minute, second, period);
    }

    public void add(Time t) {
        second += t.second;
        minute += t.minute + second / 60;
        hour += t.hour + minute / 60;

        second %= 60;
        minute %= 60;
        hour %= 24;
    }
}
