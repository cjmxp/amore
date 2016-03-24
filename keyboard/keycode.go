package keyboard

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Key corresponds to a keyboard key.
type Key sdl.Keycode

const (
	KeyUnknown            Key = Key(sdl.K_UNKNOWN)
	KeyReturn             Key = Key(sdl.K_RETURN)
	KeyEscape             Key = Key(sdl.K_ESCAPE)
	KeyBackSpace          Key = Key(sdl.K_BACKSPACE)
	KeyTab                Key = Key(sdl.K_TAB)
	KeySpace              Key = Key(sdl.K_SPACE)
	KeyExclaim            Key = Key(sdl.K_EXCLAIM)
	KeyDblQuote           Key = Key(sdl.K_QUOTEDBL)
	KeyHash               Key = Key(sdl.K_HASH)
	KeyPercent            Key = Key(sdl.K_PERCENT)
	KeyDollar             Key = Key(sdl.K_DOLLAR)
	KeyAmpersand          Key = Key(sdl.K_AMPERSAND)
	KeyQuote              Key = Key(sdl.K_QUOTE)
	KeyLeftparen          Key = Key(sdl.K_LEFTPAREN)
	KeyRightparen         Key = Key(sdl.K_RIGHTPAREN)
	KeyAsterisk           Key = Key(sdl.K_ASTERISK)
	KeyPlus               Key = Key(sdl.K_PLUS)
	KeyComma              Key = Key(sdl.K_COMMA)
	KeyMinus              Key = Key(sdl.K_MINUS)
	KeyPeriod             Key = Key(sdl.K_PERIOD)
	KeySlash              Key = Key(sdl.K_SLASH)
	Key0                  Key = Key(sdl.K_0)
	Key1                  Key = Key(sdl.K_1)
	Key2                  Key = Key(sdl.K_2)
	Key3                  Key = Key(sdl.K_3)
	Key4                  Key = Key(sdl.K_4)
	Key5                  Key = Key(sdl.K_5)
	Key6                  Key = Key(sdl.K_6)
	Key7                  Key = Key(sdl.K_7)
	Key8                  Key = Key(sdl.K_8)
	Key9                  Key = Key(sdl.K_9)
	KeyColon              Key = Key(sdl.K_COLON)
	KeySemicolon          Key = Key(sdl.K_SEMICOLON)
	KeyLess               Key = Key(sdl.K_LESS)
	KeyEquals             Key = Key(sdl.K_EQUALS)
	KeyGreater            Key = Key(sdl.K_GREATER)
	KeyQuestion           Key = Key(sdl.K_QUESTION)
	KeyAt                 Key = Key(sdl.K_AT)
	KeyLeftbracket        Key = Key(sdl.K_LEFTBRACKET)
	KeyBackslash          Key = Key(sdl.K_BACKSLASH)
	KeyRightbracket       Key = Key(sdl.K_RIGHTBRACKET)
	KeyCaret              Key = Key(sdl.K_CARET)
	KeyUnderscore         Key = Key(sdl.K_UNDERSCORE)
	KeyBackquote          Key = Key(sdl.K_BACKQUOTE)
	KeyA                  Key = Key(sdl.K_a)
	KeyB                  Key = Key(sdl.K_b)
	KeyC                  Key = Key(sdl.K_c)
	KeyD                  Key = Key(sdl.K_d)
	KeyE                  Key = Key(sdl.K_e)
	KeyF                  Key = Key(sdl.K_f)
	KeyG                  Key = Key(sdl.K_g)
	KeyH                  Key = Key(sdl.K_h)
	KeyI                  Key = Key(sdl.K_i)
	KeyJ                  Key = Key(sdl.K_j)
	KeyK                  Key = Key(sdl.K_k)
	KeyL                  Key = Key(sdl.K_l)
	KeyM                  Key = Key(sdl.K_m)
	KeyN                  Key = Key(sdl.K_n)
	KeyO                  Key = Key(sdl.K_o)
	KeyP                  Key = Key(sdl.K_p)
	KeyQ                  Key = Key(sdl.K_q)
	KeyR                  Key = Key(sdl.K_r)
	KeyS                  Key = Key(sdl.K_s)
	KeyT                  Key = Key(sdl.K_t)
	KeyU                  Key = Key(sdl.K_u)
	KeyV                  Key = Key(sdl.K_v)
	KeyW                  Key = Key(sdl.K_w)
	KeyX                  Key = Key(sdl.K_x)
	KeyY                  Key = Key(sdl.K_y)
	KeyZ                  Key = Key(sdl.K_z)
	KeyCapslock           Key = Key(sdl.K_CAPSLOCK)
	KeyF1                 Key = Key(sdl.K_F1)
	KeyF2                 Key = Key(sdl.K_F2)
	KeyF3                 Key = Key(sdl.K_F3)
	KeyF4                 Key = Key(sdl.K_F4)
	KeyF5                 Key = Key(sdl.K_F5)
	KeyF6                 Key = Key(sdl.K_F6)
	KeyF7                 Key = Key(sdl.K_F7)
	KeyF8                 Key = Key(sdl.K_F8)
	KeyF9                 Key = Key(sdl.K_F9)
	KeyF10                Key = Key(sdl.K_F10)
	KeyF11                Key = Key(sdl.K_F11)
	KeyF12                Key = Key(sdl.K_F12)
	KeyPrintscreen        Key = Key(sdl.K_PRINTSCREEN)
	KeyScrolllock         Key = Key(sdl.K_SCROLLLOCK)
	KeyPause              Key = Key(sdl.K_PAUSE)
	KeyInsert             Key = Key(sdl.K_INSERT)
	KeyHome               Key = Key(sdl.K_HOME)
	KeyPageup             Key = Key(sdl.K_PAGEUP)
	KeyDelete             Key = Key(sdl.K_DELETE)
	KeyEnd                Key = Key(sdl.K_END)
	KeyPagedown           Key = Key(sdl.K_PAGEDOWN)
	KeyRight              Key = Key(sdl.K_RIGHT)
	KeyLeft               Key = Key(sdl.K_LEFT)
	KeyDown               Key = Key(sdl.K_DOWN)
	KeyUp                 Key = Key(sdl.K_UP)
	KeyNumlockclear       Key = Key(sdl.K_NUMLOCKCLEAR)
	KeyApplication        Key = Key(sdl.K_APPLICATION)
	KeyPower              Key = Key(sdl.K_POWER)
	KeyF13                Key = Key(sdl.K_F13)
	KeyF14                Key = Key(sdl.K_F14)
	KeyF15                Key = Key(sdl.K_F15)
	KeyF16                Key = Key(sdl.K_F16)
	KeyF17                Key = Key(sdl.K_F17)
	KeyF18                Key = Key(sdl.K_F18)
	KeyF19                Key = Key(sdl.K_F19)
	KeyF20                Key = Key(sdl.K_F20)
	KeyF21                Key = Key(sdl.K_F21)
	KeyF22                Key = Key(sdl.K_F22)
	KeyF23                Key = Key(sdl.K_F23)
	KeyF24                Key = Key(sdl.K_F24)
	KeyExecute            Key = Key(sdl.K_EXECUTE)
	KeyHelp               Key = Key(sdl.K_HELP)
	KeyMenu               Key = Key(sdl.K_MENU)
	KeySelect             Key = Key(sdl.K_SELECT)
	KeyStop               Key = Key(sdl.K_STOP)
	KeyAgain              Key = Key(sdl.K_AGAIN)
	KeyUndo               Key = Key(sdl.K_UNDO)
	KeyCut                Key = Key(sdl.K_CUT)
	KeyCopy               Key = Key(sdl.K_COPY)
	KeyPaste              Key = Key(sdl.K_PASTE)
	KeyFind               Key = Key(sdl.K_FIND)
	KeyMute               Key = Key(sdl.K_MUTE)
	KeyVolumeup           Key = Key(sdl.K_VOLUMEUP)
	KeyVolumedown         Key = Key(sdl.K_VOLUMEDOWN)
	KeyCancel             Key = Key(sdl.K_CANCEL)
	KeyClear              Key = Key(sdl.K_CLEAR)
	KeyPrior              Key = Key(sdl.K_PRIOR)
	KeyReturn2            Key = Key(sdl.K_RETURN2)
	KeySeparator          Key = Key(sdl.K_SEPARATOR)
	KeyOut                Key = Key(sdl.K_OUT)
	KeyOper               Key = Key(sdl.K_OPER)
	KeyClearagain         Key = Key(sdl.K_CLEARAGAIN)
	KeyCrsel              Key = Key(sdl.K_CRSEL)
	KeyExsel              Key = Key(sdl.K_EXSEL)
	KeyThousandsSeparator Key = Key(sdl.K_THOUSANDSSEPARATOR)
	KeyDecimalSeparator   Key = Key(sdl.K_DECIMALSEPARATOR)
	KeyCurrencyUnit       Key = Key(sdl.K_CURRENCYUNIT)
	KeyCurrencySubUnit    Key = Key(sdl.K_CURRENCYSUBUNIT)
	KeyLCtrl              Key = Key(sdl.K_LCTRL)
	KeyLShift             Key = Key(sdl.K_LSHIFT)
	KeyLAlt               Key = Key(sdl.K_LALT)
	KeyLGui               Key = Key(sdl.K_LGUI)
	KeyRCtrl              Key = Key(sdl.K_RCTRL)
	KeyRShift             Key = Key(sdl.K_RSHIFT)
	KeyRAlt               Key = Key(sdl.K_RALT)
	KeyRGUI               Key = Key(sdl.K_RGUI)
	KeyMode               Key = Key(sdl.K_MODE)
	KeyAudioNext          Key = Key(sdl.K_AUDIONEXT)
	KeyAudioPrev          Key = Key(sdl.K_AUDIOPREV)
	KeyAudioStop          Key = Key(sdl.K_AUDIOSTOP)
	KeyAudioPlay          Key = Key(sdl.K_AUDIOPLAY)
	KeyAudioMute          Key = Key(sdl.K_AUDIOMUTE)
	KeyMediaSelect        Key = Key(sdl.K_MEDIASELECT)
)

// ModifierKey corresponds to a modifier key.
type ModifierKey sdl.Keymod

//// Modifier keys
const (
	ModNone     ModifierKey = ModifierKey(sdl.KMOD_NONE)
	ModLShift   ModifierKey = ModifierKey(sdl.KMOD_LSHIFT)
	ModRShift   ModifierKey = ModifierKey(sdl.KMOD_RSHIFT)
	ModLCtrl    ModifierKey = ModifierKey(sdl.KMOD_LCTRL)
	ModRCtrl    ModifierKey = ModifierKey(sdl.KMOD_RCTRL)
	ModLAlt     ModifierKey = ModifierKey(sdl.KMOD_LALT)
	ModRAlt     ModifierKey = ModifierKey(sdl.KMOD_RALT)
	ModLGui     ModifierKey = ModifierKey(sdl.KMOD_LGUI)
	ModRGui     ModifierKey = ModifierKey(sdl.KMOD_RGUI)
	ModNum      ModifierKey = ModifierKey(sdl.KMOD_NUM)
	ModCaps     ModifierKey = ModifierKey(sdl.KMOD_CAPS)
	ModMode     ModifierKey = ModifierKey(sdl.KMOD_MODE)
	ModCtrl     ModifierKey = ModifierKey(sdl.KMOD_CTRL)
	ModShift    ModifierKey = ModifierKey(sdl.KMOD_SHIFT)
	ModAlt      ModifierKey = ModifierKey(sdl.KMOD_ALT)
	ModGui      ModifierKey = ModifierKey(sdl.KMOD_GUI)
	ModReserved ModifierKey = ModifierKey(sdl.KMOD_RESERVED)
)
