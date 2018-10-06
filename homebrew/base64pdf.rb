require 'rbconfig'
class Base64pdf < Formula
  desc ""
  homepage "https://github.com/vimukthi-git/base64pdf"
  version "0.0.1"

  if Hardware::CPU.is_64_bit?
    case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
      :windows
    when /darwin|mac os/
      url "https://github.com/vimukthi-git/base64pdf/releases/download/v0.0.1/base64pdf_v0.0.1_darwin_amd64.zip"
      sha256 "2dcd41a6df212da2f11a4ba0e4089e2070f52c62fcaa9de286f3ada393c2af7c"
    when /linux/
      url "https://github.com/vimukthi-git/base64pdf/releases/download/v0.0.1/base64pdf_v0.0.1_linux_amd64.tar.gz"
      sha256 "28ab700401920b9db30fe7d6721c5103a73c769802d2591dee8df015540917e9"
    when /solaris|bsd/
      :unix
    else
      :unknown
    end
  else
    case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
      :windows
    when /darwin|mac os/
      url "https://github.com/vimukthi-git/base64pdf/releases/download/v0.0.1/base64pdf_v0.0.1_darwin_386.zip"
      sha256 "889111c5dc01b51b0e25a0221409a2ddcf9d6f101ca7d36482932c3411c74e7a"
    when /linux/
      url "https://github.com/vimukthi-git/base64pdf/releases/download/v0.0.1/base64pdf_v0.0.1_linux_386.tar.gz"
      sha256 "44d950c7f0e0a7bd281ca529b39bda85140857d94ce85f1020fad39b73534644"
    when /solaris|bsd/
      :unix
    else
      :unknown
    end
  end

  def install
    bin.install "base64pdf"
  end

  test do
    system "base64pdf"
  end

end
